import React from 'react'

import RateLabel from './RateLabel'
import FetchingComponent from './FetchingComponent'

import * as totals from '../backends/totals'

export default class MiniDashboard extends FetchingComponent {
    constructor(props) {
        super(props, totals.getRates)
    }

    renderData({rate_5sec}) {
        return (
            <div style={{position: "absolute", left: 0, right: 0, top: 0, bottom: 0}}>
                <RateLabel title="IN"
                           className="rate-in"
                           style={{left: 0, top: 0, height: "100%", width: "50%"}}
                           rate={rate_5sec.bytes_in}/>
                <RateLabel title="OUT"
                           className="rate-out"
                           style={{right: 0, top: 0, height: "100%", width: "50%"}}
                           rate={rate_5sec.bytes_out}/>
            </div>
        )
    }
}
