import React from 'react'
import ReactDOM from 'react-dom'
import {Router, Route} from "react-router"
import history from './utils/history'

import Overview from './components/Overview'

const routes = (
    <Router history={history}>
        <Route path="/" component={Overview}/>
    </Router>
)

const appRoot = window.document.getElementById("app")
ReactDOM.render(routes, appRoot)