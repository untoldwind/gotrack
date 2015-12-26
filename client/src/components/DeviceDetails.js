import React from 'react'

import DeviceOverview from './DeviceOverview'
import Fetcher from './Fetcher'
import * as devices from '../backends/devices'

export default class DeviceDetails extends React.Component {
    constructor(props) {
        super(props)

        this.fetchDetails = this.fetchDetails.bind(this)
    }

    render() {
        return (
            <Fetcher fetcher={this.fetchDetails}>
                <DeviceOverview/>
            </Fetcher>
        )
    }

    fetchDetails() {
        const {deviceIp} = this.props.params

        return devices.getDeviceDetails(deviceIp)
    }
}