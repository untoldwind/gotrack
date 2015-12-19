import React from 'react'
import AutoscaleLabel from './AutoscaleLabel'

import shallowEqual from '../utils/shallowEqual'

export default class RateLabel extends React.Component {
    static propTypes = {
        title: React.PropTypes.string.isRequired,
        className: React.PropTypes.string,
        rate: React.PropTypes.number.isRequired,
        style: React.PropTypes.object.isRequired
    }

    static defaultProps = {
        style: {
            top: "0px",
            bottom: "0px",
            left: "0px",
            right: "0px"
        }
    }

    shouldComponentUpdate(nextProps, nextState) {
        return !shallowEqual(this.props, nextProps, true) || !shallowEqual(this.state, nextState)
    }

    render() {
        const rate = this.props.rate
        var value = rate.toString()
        var unit = "bytes/s"

        if (rate >= 100 * 1024) {
            value = (rate / 1024.0 / 1024.0).toFixed(2)
            unit = "MB/s"
        } else if (rate >= 100) {
            value = (rate / 1024.0).toFixed(2)
            unit = "kB/s"
        }

        return (
            <div className={this.props.className} style={this.props.style}>
                <AutoscaleLabel innerClassName="rate-title"
                                style={{left: 0, top: 0, right: 0, height: "20%"}}
                                text={this.props.title} />
                <AutoscaleLabel innerClassName="rate-value"
                                style={{left: 0, top: "20%", right: 0, bottom: "20%"}}
                                text={value} />
                <AutoscaleLabel innerClassName="rate-unit"
                                style={{left: "50%", height: "20%", right: 0, bottom: 0}}
                                text={unit} />
            </div>
        )
    }
}
