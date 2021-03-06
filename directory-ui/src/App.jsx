import React, { Component } from 'react'
import Navigation from './components/Navigation'
import Directory from './Directory'
import Doc from './Documentation'

import {
    BrowserRouter as Router,
    Route
  } from 'react-router-dom'

class App extends Component {
    render() {
        return (
            <div className="App">
                <Router>
                    <div>
                        <Navigation />
                        <Route exact path="/" component={Directory} />
                        <Route path="/documentation/:organization_name/:service_name" component={Doc} />
                    </div>
                </Router>
            </div>
        )
    }
}

export default App;
