import React from 'react'

import OverviewRates from './OverviewRates'
import ConnectionList from './ConnectionList'
import SpanGraph from './SpanGraph'
import FetchingComponent from './FetchingComponent'

import * as devices from '../backends/devices'

export default class DeviceOverview extends FetchingComponent {
    constructor(props) {
        super(props)

        this.fetcher = this.fetchDetails.bind(this)
        this.fetchSpan = this.fetchSpan.bind(this)
    }

    renderData(data) {
        return (
            <div className="container">
                <div className="row">
                    <div className="col-xs-12">
                        <div className="panel panel-default">
                            <div className="panel-body">
                                <h2>{data.name}</h2>
                                {data.ip_address} (MAC: {data.mac_address})
                            </div>
                        </div>
                    </div>
                </div>
                <div className="row">
                    <div className="col-xs-12">
                        <div className="panel panel-default">
                            <div className="panel-body">
                                <SpanGraph fetcher={this.fetchSpan}/>
                            </div>
                        </div>
                    </div>
                </div>
                <OverviewRates data={data}/>
                <div className="row">
                    <div className="col-xs-12">
                        <div className="panel panel-default">
                            <div className="panel-body">
                                <ConnectionList data={data.connections}/>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }

    fetchDetails() {
        const {deviceIp} = this.props.params

        return devices.getDeviceDetails(deviceIp)
    }

    fetchSpan() {
        const {deviceIp} = this.props.params

        return devices.getDeviceSpan(deviceIp)
    }
}