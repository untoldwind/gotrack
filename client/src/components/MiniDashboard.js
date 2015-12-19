import React from 'react'
import Transmit from "react-transmit"

import * as totals from '../backends/totals'

import RateLabel from './RateLabel'

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
        return (
            <div style={{position: "absolute", left: 0, right: 0, top: 0, bottom: 0}}>
                <RateLabel title="IN"
                           className="rate-in"
                           style={{left: 0, top: 0, height: "100%", width: "50%"}}
                           rate={this.props.totalRates.rate_5sec.bytes_in}/>
                <RateLabel title="OUT"
                           className="rate-out"
                           style={{right: 0, top: 0, height: "100%", width: "50%"}}
                           rate={this.props.totalRates.rate_5sec.bytes_out}/>
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
