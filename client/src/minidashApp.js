import React from 'react'
import ReactDOM from 'react-dom'
import Fetcher from './components/Fetcher'
import MiniDashboard from './components/MiniDashboard'

import * as totals from './backends/totals'

class MinidashApp extends React.Component {
    render() {
        return (
            <Fetcher fetcher={totals.getRates}>
                <MiniDashboard/>
            </Fetcher>
        )
    }
}

const appRoot = window.document.getElementById("app")
ReactDOM.render(<MinidashApp/>, appRoot)
