import React from 'react'
import Transmit from "react-transmit"

import * as totals from '../backends/totals'

import BigLabel from './BigLabel'

class MiniDashboard extends React.Component {
    constructor(props) {
        super(props)
    }

    componentDidMount() {
        this.timer = window.setInterval(() => {
            this.props.transmit.forceFetch({})
        }, 2000)
    }

    componentWillUnmount() {
        window.clearTimeout(this.timer)
    }

    render() {
        console.log(this.props)
        return (
            <div style={{position: "absolute", left: 0, right: 0, top: 0, bottom: 0}}>
                <BigLabel text={this.props.totalRates.rate_5sec.bytes_in.toString()}
                          style={{ position: "absolute", left: 0, top: 0, height: "100%", width: "50%"}}/>
                <BigLabel text={this.props.totalRates.rate_5sec.bytes_out.toString()}
                          style={{ position: "absolute", left: "50%", top: 0, height: "100%", width: "50%"}}/>
            </div>
        )
    }
}

export default Transmit.createContainer(MiniDashboard, {
    initialVariables: {},
    fragments: {
        totalRates() {
            return totals.getRates()
        }
    }
})
