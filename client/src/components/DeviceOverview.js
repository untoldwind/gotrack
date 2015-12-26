import React from 'react'

import OverviewRates from './OverviewRates'
import ConnectionList from './ConnectionList'
import Fetcher from './Fetcher'
import SpanGraph from './SpanGraph'

import * as devices from '../backends/devices'

export default class DeviceOverview extends React.Component {
    constructor(props) {
        super(props)

        this.fetchSpan = this.fetchSpan.bind(this)
    }

    render() {
        return (
            <div className="container">
                <div className="row">
                    <div className="col-xs-12">
                        <div className="panel panel-default">
                            <div className="panel-body">
                                <h2>{this.props.data.name}</h2>
                                {this.props.data.ip_address} (MAC: {this.props.data.mac_address})
                            </div>
                        </div>
                    </div>
                </div>
                <div className="row">
                    <div className="col-xs-12">
                        <div className="panel panel-default">
                            <div className="panel-body">
                                <Fetcher fetcher={this.fetchSpan}>
                                    <SpanGraph/>
                                </Fetcher>
                            </div>
                        </div>
                    </div>
                </div>
                <OverviewRates data={this.props.data}/>
                <div className="row">
                    <div className="col-xs-12">
                        <div className="panel panel-default">
                            <div className="panel-body">
                                <ConnectionList data={this.props.data.connections}/>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }

    fetchSpan() {
        return devices.getDeviceSpan(this.props.data.ip_address)
    }
}