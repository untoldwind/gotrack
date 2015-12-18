import React from 'react'
import ReactDOM from "react-dom"

import {throttle} from '../utils/dash'

export default class Rescale extends React.Component {
    static propTypes = {
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
        console.log(this.state)
        const outerStyle = Object.assign({}, this.props.style)
        const child = React.cloneElement(React.Children.only(this.props.children), {
            fontSize: this.state.fontSize,
            totalHeight: this.state.height,
            ref: "inner"
        })
        return (
            <div className="rescale" style={outerStyle} ref="outer">
                {child}
            </div>
        )
    }

    rescale() {
        const outer = this.refs.outer
        const inner = ReactDOM.findDOMNode(this.refs.inner)

        if (outer.offsetHeight !== this.state.height) {
            this.setState({
                height: outer.offsetHeight
            })
        } else {
            console.log(outer.offsetWidth)
            console.log(inner.offsetWidth)
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
