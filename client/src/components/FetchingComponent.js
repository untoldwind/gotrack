import React from 'react'

export default class FetcherComponent extends React.Component {
    constructor(props, fetcher) {
        super(props)

        this.fetcher = fetcher
        this.state = { fetchedData: null }

        if (this.constructor === FetcherComponent) {
            throw new TypeError('Abstract class "FetcherComponent" cannot be instantiated directly.');
        }
    }

    componentDidMount() {
        this.fetch()
        this.timer = window.setInterval(this.fetch.bind(this), 2000)
    }

    componentWillUnmount() {
        window.clearTimeout(this.timer)
    }

    render() {
        if (!this.state.fetchedData) {
            return (
                <div ref="node"></div>
            )
        } else {
            return this.renderData(this.state.fetchedData)
        }
    }

    fetch() {
        this.fetcher().then(
            data => this.setState({fetchedData: data})
        )
    }
}