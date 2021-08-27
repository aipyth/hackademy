import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  Redirect
} from "react-router-dom";
import './App.css';

import { Home } from './components/Home';
import { Login } from './components/Login';
import { Forecast } from './components/Forecast';
import Storage from './Storage';

interface AppProps {
};
interface AppState {
    logged: boolean,
};
class App extends React.Component<AppProps, AppState> {
    constructor(props: AppProps) {
        super(props);
        this.state = {
            logged: Storage.hasUser(),
        };
        this.loginHandler = this.loginHandler.bind(this);
    }

    loginHandler(logged: boolean): void {
        this.setState({ logged: logged });
    }

    renderLinks() {
        return (
            <nav className="hidden md:block text-2xl space-x-4 md:mr-8">
                <span><Link to="/">Home</Link></span>
                { this.state.logged ? null : <span><Link to="/login">Login</Link></span> }
                { this.state.logged ? <span><Link to="/forecast">Forecast</Link></span> : null }
            </nav>
        );
    }

    render() {
        return (
        <Router>
            <header className="w-full top-0 fixed flex justify-center md:justify-between items-center h-24 bg-yellow-500 text-black">
                <span className="logo md:ml-8">Landing</span>
                {this.renderLinks()}
            </header>

            <div className="mt-24">
                <Switch>
                    <Route exact path="/">
                        <Home />
                    </Route>
                    <Route path="/login">
                        {this.state.logged ? <Redirect to="/" /> : <Login loginHandler={this.loginHandler} /> }
                    </Route>
                    <Route path="/forecast">
                        {this.state.logged ? <Forecast /> : <Redirect to="/login" /> }
                    </Route>
                    <Route path="*">
                        <Redirect to="/" />
                    </Route>
                </Switch>
            </div>

        </Router>
        );
    }
};

export default App;
