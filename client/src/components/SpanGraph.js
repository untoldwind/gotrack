import React from 'react'
import ReactART from 'react-art'
import ReactDOM from 'react-dom'
import FetchingComponent from './FetchingComponent'
import Rectangle from 'react-art/lib/Rectangle.art'

import {throttle} from '../utils/dash'
import {formatTotalString} from './formats'

const {Surface, Group, Shape, Text, LinearGradient, Path} = ReactART

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

        const upperRect = Path().moveTo(0, height/2).lineTo(width, height/2).lineTo(width, 0).lineTo(0, 0)
        const lowerRect = Path().moveTo(0, height/2).lineTo(width, height/2).lineTo(width, height).lineTo(0, height)
        const inLabelX = width * 0.1
        const inLabelY =  height * 0.25
        const maxInLine = Path().moveTo(inLabelX, inLabelY).lineTo(inLabelX + inLabelY - 5, 5).lineTo(width, 5)
        const outLabelX = width * 0.1
        const outLabelY = height * 0.75
        const maxOutLine = Path().moveTo(outLabelX, outLabelY).lineTo(outLabelX + (graphHeight - outLabelY) - 5, graphHeight).lineTo(width, graphHeight)
        const baseLine = Path().moveTo(0, height/2).lineTo(width, height/2)

        return (
            <Surface ref="node" width={width} height={height} style={{backgroundColor: "black"}}>
                <Group>
                    <Shape fill={new LinearGradient(['#040', '#000'], 0, height/2, 0, 0)} d={upperRect}/>
                    <Shape fill={new LinearGradient(['#400', '#000'], 0, height/2, 0, height)} d={lowerRect}/>
                    <Shape fill="#0f0" d={inGraph}/>
                    <Shape fill="#f00" d={outGraph}/>
                    <Shape stroke="#aaa" d={maxInLine}/>
                    <Shape stroke="#aaa" d={maxOutLine}/>
                    <Rectangle fill="#222" x={inLabelX - 50} y={inLabelY - 10} width={100} height={20} radius={5}/>
                    <Rectangle fill="#222" x={outLabelX - 50} y={outLabelY - 10} width={100} height={20} radius={5}/>
                    <Text fill="#eee" font='16px "Arial"' alignment='middle' x={inLabelX} y={inLabelY - 8}>{formatTotalString(max.bytes_in)}</Text>
                    <Text fill="#eee" font='16px "Arial"' alignment='middle' x={outLabelX} y={outLabelY - 8}>{formatTotalString(max.bytes_out)}</Text>
                    <Shape stroke="#fff" d={baseLine}/>
                </Group>
            </Surface>
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