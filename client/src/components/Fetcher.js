import React from 'react'

export default class Fetcher extends React.Component {
    static propTypes = {
        fetcher: React.PropTypes.func.isRequired
    }

    constructor(props) {
        super(props)

        this.state = { fetchedData: null }
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
                <div></div>
            )
        } else {
            const child = React.Children.only(this.props.children)

            return React.cloneElement(child, {data: this.state.fetchedData})
        }
    }

    fetch() {
        this.props.fetcher().then(
            data => this.setState({fetchedData: data})
        )
    }
}