import React from 'react'

export default class DeviceList extends React.Component {
    constructor(props) {
        super(props)

        this.state = {sortField: "name", sortAsc: true}
    }

    render() {
        const sorted = this.props.data.sort(this.sortFunction())
        const rows = sorted.map((device) => {
            return (
                <tr key={device.ip_address}>
                    <td>{device.name}</td>
                    <td>{device.ip_address}</td>
                    <td>{device.mac_address}</td>
                    <td>{device.connection_count}</td>
                    <td className="rate-in">{device.rate_5sec.bytes_in}</td>
                    <td className="rate-out">{device.rate_5sec.bytes_out}</td>
                </tr>
            )
        })
        return (
            <table className="table table-hover">
                <thead>
                <tr>
                    <th onClick={this.switchSort("name")}>Name {this.sortMarker("name")}</th>
                    <th onClick={this.switchSort("ip_address")}>IP {this.sortMarker("ip_address")}</th>
                    <th onClick={this.switchSort("mac_address")}>MAC {this.sortMarker("mac_address")}</th>
                    <th>Connections</th>
                    <th className="rate-in">In (b/s)</th>
                    <th className="rate-out">Out (b/s)</th>
                </tr>
                </thead>
                <tbody>
                {rows}
                </tbody>
            </table>
        )
    }

    sortFunction() {
        const field = this.state.sortField

        if (this.state.sortAsc) {
            return (a, b) => a[field] > b[field]
        } else {
            return (a, b) => b[field] > a[field]
        }
    }

    sortMarker(field) {
        if (field === this.state.sortField) {
            if (this.state.sortAsc) {
                return <span className="glyphicon glyphicon-chevron-down" ariaHidden="true"></span>
            } else {
                return <span className="glyphicon glyphicon-chevron-up" ariaHidden="true"></span>
            }
        }
        return <span></span>
    }

    switchSort(field) {
        return () => {
            if (field === this.state.sortField) {
                this.setState({sortAsc: !this.state.sortAsc})
            } else {
                this.setState({sortField: field, sortAsc: true})
            }
        }
    }
}
