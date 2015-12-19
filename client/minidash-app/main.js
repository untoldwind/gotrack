'use strict';

const argv = require('minimist')(process.argv);
const electron = require('electron');
const app = electron.app;
const BrowserWindow = electron.BrowserWindow;

var mainWindow = null;

var host = argv.host || 'localhost:8080'

app.on('window-all-closed', function() {
    if (process.platform !== 'darwin') app.quit();
});

app.on('ready', function() {
    mainWindow = new BrowserWindow({});
    mainWindow.setFullScreen(true);
    mainWindow.loadURL('http://' + host + '/minidash.html');

    mainWindow.on('closed', function() {
        mainWindow = null;
    });
});