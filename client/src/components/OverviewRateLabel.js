import React from 'react'

import {formatRate} from './formats'

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
        const {value, unit} = formatRate(this.props.rate)

        return (
            <div className="panel panel-default">
                <div className="panel-body">
                    <div className={"col-xs-2 rate-title " + this.props.className}>{this.props.title}</div>
                    <div className={"col-xs-7 rate-value " + this.props.className}>{value}</div>
                    <div className={"col-xs-3 rate-unit " + this.props.className}>{unit}</div>
                </div>
            </div>
        )
    }
}