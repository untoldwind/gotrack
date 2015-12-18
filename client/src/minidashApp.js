import React from 'react'
import Transmit from "react-transmit"

import MiniDashboard from './components/MiniDashboard'

class MinidashApp extends React.Component {
    render() {
        return (
            <MiniDashboard/>
        )
    }
}

const appRoot = window.document.getElementById("app")
Transmit.render(MinidashApp, {}, appRoot)
