import React from 'react'

import DeviceList from './DeviceList'
import OverviewRates from './OverviewRates'
import SpanGraph from './SpanGraph'
import * as totals from '../backends/totals'

export default class Overview extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="row">
                    <div className="col-xs-12">
                        <div className="panel panel-default">
                            <div className="panel-body">
                                <SpanGraph fetcher={totals.getSpan}/>
                            </div>
                        </div>
                    </div>
                </div>
                <OverviewRates/>
                <div className="row">
                    <div className="col-xs-12">
                        <div className="panel panel-default">
                            <div className="panel-body">
                                <DeviceList/>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}