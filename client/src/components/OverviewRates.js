import React from 'react'

import OverviewRateLabel from './OverviewRateLabel'

import shallowEqual from '../utils/shallowEqual'

export default class OverviewRates extends React.Component {
    constructor(props) {
        super(props)
    }

    shouldComponentUpdate(nextProps, nextState) {
        return !shallowEqual(this.props, nextProps, true) || !shallowEqual(this.state, nextState)
    }

    render() {
        return (
            <div className="row">
                <div className="col-md-6">
                    <OverviewRateLabel className="rate-in" title="In" rate={this.props.data.rate_5sec.bytes_in}/>
                </div>
                <div className="col-md-6">
                    <OverviewRateLabel className="rate-out" title="Out" rate={this.props.data.rate_5sec.bytes_out}/>
                </div>
            </div>
        )
    }
}