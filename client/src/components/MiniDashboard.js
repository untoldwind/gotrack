import React from 'react'

import RateLabel from './RateLabel'

export default class MiniDashboard extends React.Component {
    constructor(props) {
        super(props)
    }

    render() {
        return (
            <div style={{position: "absolute", left: 0, right: 0, top: 0, bottom: 0}}>
                <RateLabel title="IN"
                           className="rate-in"
                           style={{left: 0, top: 0, height: "100%", width: "50%"}}
                           rate={this.props.data.rate_5sec.bytes_in}/>
                <RateLabel title="OUT"
                           className="rate-out"
                           style={{right: 0, top: 0, height: "100%", width: "50%"}}
                           rate={this.props.data.rate_5sec.bytes_out}/>
            </div>
        )
    }
}
