import React from 'react'

import shallowEqual from '../utils/shallowEqual'
import {throttle} from '../utils/dash'

export default class BigLabel extends React.Component {
    static propTypes = {
        text: React.PropTypes.string.isRequired,
        style: React.PropTypes.object.isRequired
    }

    static defaultProps = {
        style: {
            position: "absolute",
            top: "0px",
            bottom: "0px",
            left: "0px",
            right: "0px"
        }
    }

    constructor(props) {
        super(props)

        this.updateDimensions = throttle(this.updateDimensions.bind(this), 500)
        this.state            = {
            fontSize: 32,
            step: 16,
            height: null,
            grow: true
        }
    }

    shouldComponentUpdate(nextProps, nextState) {
        return !shallowEqual(this.props, nextProps, true) || !shallowEqual(this.state, nextState)
    }

    componentWillReceiveProps() {
        this.setState({
            fontSize: 32,
            step: 16,
            height: null,
            grow: true
        })
    }

    componentDidMount() {
        window.addEventListener("resize", this.updateDimensions)
        this.rescale()
    }

    componentWillUnmount() {
        window.removeEventListener("resize", this.updateDimensions)
    }

    componentDidUpdate() {
        this.rescale()
    }

    render() {
        const outerStyle = Object.assign({}, this.props.style)
        if (this.state.height) {
            outerStyle['lineHeight'] = this.state.height + "px"
        }
        const innerStyle = {
            fontSize: this.state.fontSize
        }
        return (
            <div className="bigLabel" style={outerStyle} ref="outer">
                <span style={innerStyle} ref="inner">{this.props.text}</span>
            </div>
        )
    }

    rescale() {
        const outer = this.refs.outer
        const inner = this.refs.inner

        if (!this.state.height) {
            this.setState({
                height: outer.offsetHeight
            })
        } else if (inner.offsetWidth > outer.offsetWidth || inner.offsetHeight > outer.offsetHeight) {
            this.decreaseSize()
        } else {
            this.increaseSize()
        }
    }


    decreaseSize() {
        const nextStep = this.state.step > 1 ? this.state.step / 2 : 0
        const step     = nextStep > 0 ? nextStep : 1
        this.setState({fontSize: this.state.fontSize - step, step: nextStep, grow: false})

    }

    increaseSize() {
        var nextStep
        if (this.state.grow) {
            nextStep = this.state.step
        } else {
            nextStep = this.state.step > 1 ? this.state.step / 2 : 0
        }
        if (nextStep > 0) {
            this.setState({fontSize: this.state.fontSize + nextStep, step: nextStep})
        }
    }

    updateDimensions() {
        this.setState({fontSize: 32, step: 16, height: null, grow: true})
    }
}