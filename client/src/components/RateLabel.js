import React from 'react'

import shallowEqual from '../utils/shallowEqual'

export default class RateLabel extends React.Component {
    static propTypes = {
        title: React.PropTypes.string.isRequired,
        className: React.PropTypes.string,
        rate: React.PropTypes.number.isRequired,
        fontSize: React.PropTypes.number,
        totalHeight: React.PropTypes.number,
    }

    static defaultProps = {
        fontSize: 32,
        totalHeight: 32
    }

    shouldComponentUpdate(nextProps, nextState) {
        return !shallowEqual(this.props, nextProps, true) || !shallowEqual(this.state, nextState)
    }

    render() {
        const style = {
            fontSize: this.props.fontSize
//            lineHeight: this.props.totalHeight + "px"
        }
        return (
            <div className={this.props.className}>
                <span className="rate-title" style={{fontSize: this.props.fontSize / 3}}>{this.props.title}</span>
                <br/>
                <span className="rate-value" style={style}>
                    {this.props.rate.toString()}
                </span>
                <br/>
                <span className="rate-unit" style={{float:"right", fontSize: this.props.fontSize / 3}}>byte/s</span>
            </div>
        )
    }
}
