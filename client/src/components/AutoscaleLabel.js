import React from 'react'
import ReactDOM from "react-dom"

import {throttle} from '../utils/dash'

export default class AutoscaleLabel extends React.Component {
    static propTypes = {
        outerClassName: React.PropTypes.string,
        innerClassName: React.PropTypes.string,
        text: React.PropTypes.string,
        style: React.PropTypes.object.isRequired
    }

    static defaultProps = {
        outerClassName: "autoscale",
        style: {
            top: "0px",
            bottom: "0px",
            left: "0px",
            right: "0px"
        }
    }

    constructor(props) {
        super(props)

        this.updateDimensions = throttle(this.rescale.bind(this), 500)
        this.state            = {
            fontSize: 32,
            height: null
        }
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
        const innerStyle = {
            fontSize: this.state.fontSize,
            lineHeight: this.state.height + "px"
        }
        return (
            <div className={this.props.outerClassName} style={outerStyle} ref="outer">
                <span className={this.props.innerClassName} style={innerStyle} ref="inner">{this.props.text}</span>
            </div>
        )
    }

    rescale() {
        const outer = this.refs.outer
        const inner = this.refs.inner

        if (outer.offsetHeight !== this.state.height) {
            this.setState({
                height: outer.offsetHeight
            })
        } else {
            const newFontsizeW = Math.floor(this.state.fontSize * outer.offsetWidth / inner.offsetWidth)
            const newFontsizeH = Math.floor(this.state.fontSize * outer.offsetHeight / inner.offsetHeight)

            if (newFontsizeW < newFontsizeH) {
                if (this.state.fontSize !== newFontsizeW) {
                    this.setState({
                        fontSize: newFontsizeW
                    })
                }
            } else if(this.state.fontSize !== newFontsizeH) {
                this.setState({
                    fontSize: newFontsizeH
                })
            }
        }
    }
}
