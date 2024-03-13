/* global pmtiles */
// import DataTile from 'ol/source/DataTile.js';
// import Map from 'ol/Map.js';
// import TileLayer from 'ol/layer/WebGLTile.js';
// import View from 'ol/View.js';
// import {useGeographic} from 'ol/proj.js';

 ol.proj.useGeographic();

const tiles = new pmtiles.PMTiles(
    'my_area.pmtiles'
);

function loadImage(src) {
    return new Promise((resolve, reject) => {
        const img = new Image();
        img.addEventListener('load', () => resolve(img));
        img.addEventListener('error', () => reject(new Error('load failed')));
        img.src = src;
    });
}

async function loader(z, x, y) {
    const response = await tiles.getZxy(z, x, y);
    const blob = new Blob([response.data]);
    const src = URL.createObjectURL(blob);
    const image = await loadImage(src);
    URL.revokeObjectURL(src);
    return image;
}

// The method used to extract elevations from the DEM.
//   red * 256 + green + blue / 256 - 32768
function elevation(xOffset, yOffset) {
    const red = ['band', 1, xOffset, yOffset];
    const green = ['band', 2, xOffset, yOffset];
    const blue = ['band', 3, xOffset, yOffset];

    // band math operates on normalized values from 0-1
    // so we scale by 255
    return [
        '+',
        ['*', 255 * 256, red],
        ['*', 255, green],
        ['*', 255 / 256, blue],
        -32768,
    ];
}

// Generates a shaded relief image given elevation data.  Uses a 3x3
// neighborhood for determining slope and aspect.
const dp = ['*', 2, ['resolution']];
const z0x = ['*', ['var', 'vert'], elevation(-1, 0)];
const z1x = ['*', ['var', 'vert'], elevation(1, 0)];
const dzdx = ['/', ['-', z1x, z0x], dp];
const z0y = ['*', ['var', 'vert'], elevation(0, -1)];
const z1y = ['*', ['var', 'vert'], elevation(0, 1)];
const dzdy = ['/', ['-', z1y, z0y], dp];
const slope = ['atan', ['sqrt', ['+', ['^', dzdx, 2], ['^', dzdy, 2]]]];
const aspect = ['clamp', ['atan', ['-', 0, dzdx], dzdy], -Math.PI, Math.PI];
const sunEl = ['*', Math.PI / 180, ['var', 'sunEl']];
const sunAz = ['*', Math.PI / 180, ['var', 'sunAz']];

const incidence = [
    '+',
    ['*', ['sin', sunEl], ['cos', slope]],
    ['*', ['cos', sunEl], ['sin', slope], ['cos', ['-', sunAz, aspect]]],
];

const variables = {};

const layer = new ol.layer.WebGLTile({
    source: new ol.source.DataTile({
        loader,
        wrapX: true,
        maxZoom: 19,
        attributions:
            "<a href='https://github.com/tilezen/joerd/blob/master/docs/attribution.md#attribution'>Tilezen Jörð</a>",
    }),
    style: {
        variables: variables,
        color: ['array', incidence, incidence, incidence, 1],
    },
});

const controlIds = ['vert', 'sunEl', 'sunAz'];
controlIds.forEach(function (id) {
    const control = document.getElementById(id);
    const output = document.getElementById(id + 'Out');
    function updateValues() {
        output.innerText = control.value;
        variables[id] = Number(control.value);
    }
    updateValues();
    control.addEventListener('input', function () {
        updateValues();
        layer.updateStyleVariables(variables);
    });
});

const map = new ol.Map({
    target: 'map',
    layers: [layer],
    view: new ol.View({
        center: [0, 0],
        zoom: 1,
    }),
});

function getElevation(data) {
    const red = data[0];
    const green = data[1];
    const blue = data[2];
    return red * 256 + green + blue / 256 - 32768;
}

function formatLocation([lon, lat]) {
    const NS = lat < 0 ? 'S' : 'N';
    const EW = lon < 0 ? 'W' : 'E';
    return `${Math.abs(lat).toFixed(1)}° ${NS}, ${Math.abs(lon).toFixed(
        1
    )}° ${EW}`;
}

const elevationOut = document.getElementById('elevationOut');
const locationOut = document.getElementById('locationOut');
function displayPixelValue(event) {
    const data = layer.getData(event.pixel);
    if (!data) {
        return;
    }
    elevationOut.innerText = getElevation(data).toLocaleString() + ' m';
    locationOut.innerText = formatLocation(event.coordinate);
}
map.on(['pointermove', 'click'], displayPixelValue);
