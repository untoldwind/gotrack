import React from 'react'
import ReactDOM from 'react-dom'
import {Router, Route} from "react-router"
import history from './utils/history'

import Overview from './components/Overview'
import DeviceDetails from './components/DeviceDetails'

const routes = (
    <Router history={history}>
        <Route path="/" component={Overview}/>
        <Route path="/devices/:deviceIp" component={DeviceDetails}/>
    </Router>
)

const appRoot = window.document.getElementById("app")
ReactDOM.render(routes, appRoot)