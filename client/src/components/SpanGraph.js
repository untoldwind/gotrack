import React from 'react'
import ReactART from 'react-art'
import ReactDOM from 'react-dom'
import FetchingComponent from './FetchingComponent'

import {throttle} from '../utils/dash'
import {formatTotalString} from './formats'

export default class SpanGraph extends FetchingComponent {
    static propTypes = {
        fetcher: React.PropTypes.func.isRequired
    }

    constructor(props) {
        super(props, props.fetcher)

        this.updateDimensions = throttle(this.rescale.bind(this), 500)
        this.state = {width: 100, height:20}
    }

    componentDidMount() {
        super.componentDidMount()
        window.addEventListener("resize", this.updateDimensions)
        this.rescale()
    }

    componentWillUnmount() {
        super.componentWillUnmount()
        window.removeEventListener("resize", this.updateDimensions)
    }

    componentDidUpdate() {
        this.rescale()
    }

    renderData({max, deltas}) {
        const {width, height} = this.state
        const graphHeight = height - 5
        const inGraph = this.createGraph(-1, "bytes_in", max, deltas)
        const outGraph = this.createGraph(1, "bytes_out", max, deltas)

        const upperRect = ReactART.Path().moveTo(0, height/2).lineTo(width, height/2).lineTo(width, 0).lineTo(0, 0)
        const lowerRect = ReactART.Path().moveTo(0, height/2).lineTo(width, height/2).lineTo(width, height).lineTo(0, height)
        const maxInLine = ReactART.Path().moveTo(width * 0.1, height * 0.25).lineTo(width * 0.1 + (height * 0.25), 5).lineTo(width, 5)
        const maxOutLine = ReactART.Path().moveTo(width * 0.1, height * 0.75).lineTo(width * 0.1 + (graphHeight - height * 0.75), graphHeight).lineTo(width, graphHeight)
        const baseLine = ReactART.Path().moveTo(0, height/2).lineTo(width, height/2)

        return (
            <ReactART.Surface ref="node" width={width} height={height} style={{backgroundColor: "black"}}>
                <ReactART.Group>
                    <ReactART.Shape fill={new ReactART.LinearGradient(['#040', '#000'], 0, height/2, 0, 0)} d={upperRect}/>
                    <ReactART.Shape fill={new ReactART.LinearGradient(['#400', '#000'], 0, height/2, 0, height)} d={lowerRect}/>
                    <ReactART.Shape fill="#0f0" d={inGraph}/>
                    <ReactART.Shape fill="#f00" d={outGraph}/>
                    <ReactART.Shape stroke="#aaa" d={maxInLine}/>
                    <ReactART.Shape stroke="#aaa" d={maxOutLine}/>
                    <ReactART.Text fill="#ddd" font='16px "Arial"' alignment='middle' x={width * 0.1} y={height * 0.25}>{formatTotalString(max.bytes_in)}</ReactART.Text>
                    <ReactART.Text fill="#ddd" font='16px "Arial"' alignment='middle' x={width * 0.1} y={height * 0.75 - 16}>{formatTotalString(max.bytes_out)}</ReactART.Text>
                    <ReactART.Shape stroke="#fff" d={baseLine}/>
                </ReactART.Group>
            </ReactART.Surface>
        )
    }

    rescale() {
        const parent = ReactDOM.findDOMNode(this.refs.node).parentNode
        const width = parent.offsetWidth - 30

        if (width !== this.state.width) {
            this.setState({width: width, height: width / 5})
        }
    }

    createGraph(dir, field, max, deltas) {
        const {width, height} = this.state
        const graph = ReactART.Path()

        if (max[field] > 0) {
            const len = deltas.length
            const graphHeight = height - 10

            graph.moveTo(0, height/2)
            deltas.forEach((delta, index) => {
                graph.
                    lineTo(index * width / len, (height + dir * delta[field]/max[field] * graphHeight)/2).
                    lineTo((index + 1) * width / len, (height + dir * delta[field]/max[field] * graphHeight)/2)
            })
            graph.lineTo(width, height/2)
        }

        return graph
    }
}