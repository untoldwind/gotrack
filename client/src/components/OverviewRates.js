import React from 'react'

import OverviewRateLabel from './OverviewRateLabel'
import FetchingComponent from './FetchingComponent'

import shallowEqual from '../utils/shallowEqual'

import * as totals from '../backends/totals'

export default class OverviewRates extends FetchingComponent {
    constructor(props) {
        super(props, totals.getRates)
    }

    shouldComponentUpdate(nextProps, nextState) {
        return !shallowEqual(this.props, nextProps, true) || !shallowEqual(this.state, nextState)
    }

    renderData({rate_5sec}) {
        return (
            <div className="row">
                <div className="col-xs-6">
                    <OverviewRateLabel className="rate-in" title="In" rate={rate_5sec.bytes_in}/>
                </div>
                <div className="col-xs-6">
                    <OverviewRateLabel className="rate-out" title="Out" rate={rate_5sec.bytes_out}/>
                </div>
            </div>
        )
    }
}