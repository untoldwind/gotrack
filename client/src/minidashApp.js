import React from 'react'
import ReactDOM from 'react-dom'
import MiniDashboard from './components/MiniDashboard'

class MinidashApp extends React.Component {
    render() {
        return (
            <MiniDashboard/>
        )
    }
}

const appRoot = window.document.getElementById("app")
ReactDOM.render(<MinidashApp/>, appRoot)
