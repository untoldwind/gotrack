import React from 'react'

export default class OverviewRateLabel extends React.Component {
    static propTypes = {
        title: React.PropTypes.string.isRequired,
        className: React.PropTypes.string,
        rate: React.PropTypes.number.isRequired,
    }

    constructor(props) {
        super(props)
    }

    render() {
        const rate = this.props.rate
        var value  = rate.toString()
        var unit   = "bytes/s"

        if (rate >= 100 * 1024) {
            value = (rate / 1024.0 / 1024.0).toFixed(2)
            unit  = "MB/s"
        } else if (rate >= 100) {
            value = (rate / 1024.0).toFixed(2)
            unit  = "kB/s"
        }

        return (
            <div className="panel panel-default">
                <div className="panel-body">
                    <div className={"col-md-3 rate-title " + this.props.className}>{this.props.title}</div>
                    <div className={"col-md-6 rate-value " + this.props.className}>{value}</div>
                    <div className={"col-md-3 rate-unit " + this.props.className}>{unit}</div>
                </div>
            </div>
        )
    }
}