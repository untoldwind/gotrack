import React from 'react'
import ReactART from 'react-art'
import ReactDOM from 'react-dom'

import {throttle} from '../utils/dash'

export default class SpanGraph extends React.Component {
    constructor(props) {
        super(props)

        this.updateDimensions = throttle(this.rescale.bind(this), 500)
        this.state = {width: 100, height:20}
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
        const {width, height} = this.state
        const {max, deltas} = this.props.data
        var inGraph = ReactART.Path()

        if (max.bytes_in > 0) {
            const maxIn = max.bytes_in
            const len = deltas.length
            const graphHeight = height - 10

            inGraph.moveTo(0, height/2)
            deltas.forEach((delta, index) => {
                inGraph.
                    lineTo(index * width / len, (height - delta.bytes_in/maxIn * graphHeight)/2).
                    lineTo((index + 1) * width / len, (height - delta.bytes_in/maxIn * graphHeight)/2)
            })
            inGraph.lineTo(width, height/2)
        }

        var outGraph = ReactART.Path()

        if (max.bytes_out > 0) {
            const maxOut = max.bytes_out
            const len = deltas.length
            const graphHeight = height - 10

            outGraph.moveTo(0, height/2)
            deltas.forEach((delta, index) => {
                outGraph.
                    lineTo(index * width / len, (height + delta.bytes_out/maxOut * graphHeight)/2).
                    lineTo((index + 1) * width / len, (height + delta.bytes_out/maxOut * graphHeight)/2)
            })
            outGraph.lineTo(width, height/2)
        }

        const upperRect = ReactART.Path().moveTo(0, height/2).lineTo(width, height/2).lineTo(width, 0).lineTo(0, 0)
        const lowerRect = ReactART.Path().moveTo(0, height/2).lineTo(width, height/2).lineTo(width, height).lineTo(0, height)
        const baseLine = ReactART.Path().moveTo(0, height/2).lineTo(width, height/2)

        return (
            <ReactART.Surface ref="surface" width={width} height={height} style={{backgroundColor: "black"}}>
                <ReactART.Group>
                    <ReactART.Shape fill={new ReactART.LinearGradient(['#400', '#000'], 0, height/2, 0, 0)} d={upperRect}/>
                    <ReactART.Shape fill={new ReactART.LinearGradient(['#040', '#000'], 0, height/2, 0, height)} d={lowerRect}/>
                    <ReactART.Shape fill="#f00" d={inGraph}/>
                    <ReactART.Shape fill="#0f0" d={outGraph}/>
                    <ReactART.Shape stroke="#fff" d={baseLine}/>
                </ReactART.Group>
            </ReactART.Surface>
        )
    }

    rescale() {
        const parent = ReactDOM.findDOMNode(this.refs.surface).parentNode
        const width = parent.offsetWidth - 30

        if (width !== this.state.width) {
            this.setState({width: width, height: width / 5})
        }
    }
}