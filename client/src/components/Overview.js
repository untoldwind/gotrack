import React from 'react'

import DeviceList from './DeviceList'
import OverviewRates from './OverviewRates'
import SpanGraph from './SpanGraph'
import Fetcher from './Fetcher'
import * as devices from '../backends/devices'
import * as totals from '../backends/totals'

export default class Overview extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="row">
                    <div className="col-xs-12">
                        <div className="panel panel-default">
                            <div className="panel-body">
                                <Fetcher fetcher={totals.getSpan}>
                                    <SpanGraph/>
                                </Fetcher>
                            </div>
                        </div>
                    </div>
                </div>
                <Fetcher fetcher={totals.getRates}>
                    <OverviewRates/>
                </Fetcher>
                <div className="row">
                    <div className="col-xs-12">
                        <div className="panel panel-default">
                            <div className="panel-body">
                                <Fetcher fetcher={devices.getDevices}>
                                    <DeviceList/>
                                </Fetcher>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}