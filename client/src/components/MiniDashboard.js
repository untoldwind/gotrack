import React from 'react'
import Transmit from "react-transmit"

import * as totals from '../backends/totals'

import Rescale from './Rescale'
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
        console.log(this.props)
        return (
            <div style={{position: "absolute", left: 0, right: 0, top: 0, bottom: 0}}>
                <Rescale style={{left: 0, top: 0, height: "100%", width: "50%"}}>
                    <RateLabel title="In" className="rate-in" rate={this.props.totalRates.rate_5sec.bytes_in}/>
                </Rescale>
                <Rescale style={{right: 0, top: 0, height: "100%", width: "50%"}}>
                    <RateLabel title="Out" className="rate-out" rate={this.props.totalRates.rate_5sec.bytes_out}/>
                </Rescale>
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
