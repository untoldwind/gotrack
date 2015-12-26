import React from 'react'

import {formatTotalString} from './formats'

export default class ConnectionList extends React.Component {
    constructor(props) {
        super(props)

        this.state = {sortField: "dst_host", sortAsc: true}
    }

    render() {
        const sorted = this.props.data.sort(this.sortFunction())
        const rows = sorted.map((connection, index) => {
            return (
                <tr key={index}>
                    <td>{connection.protocol}</td>
                    <td>{connection.src_port}</td>
                    <td>{connection.dst_host}</td>
                    <td>{connection.dst_port}</td>
                    <td className="rate-in">{formatTotalString(connection.totals.bytes_in)}</td>
                    <td className="rate-out">{formatTotalString(connection.totals.bytes_out)}</td>
                </tr>
            )
        })
        return (
            <table className="table table-hover">
                <thead>
                <tr>
                    <th>Protocol</th>
                    <th onClick={this.switchSort("src_port")}>Src port {this.sortMarker("src_port")}</th>
                    <th onClick={this.switchSort("dst_host")}>Dst host {this.sortMarker("dst_host")}</th>
                    <th>Dst port</th>
                    <th className="rate-in">In</th>
                    <th className="rate-out">Out</th>
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
            return (a, b) => a[field] < b[field] ? -1 : 1
        } else {
            return (a, b) => b[field] < a[field] ? -1 : 1
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