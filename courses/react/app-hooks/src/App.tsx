import React, { useState } from 'react';
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

export default function App() {
    const [logged, setLogged] = useState(Storage.hasUser());

    const loginHandler = (logged: boolean): void => {
        setLogged(logged);
    }

    const renderLinks = () => {
        return (
            <nav className="hidden md:block text-2xl space-x-4 md:mr-8">
                <span><Link to="/">Home</Link></span>
                { logged ? null : <span><Link to="/login">Login</Link></span> }
                { logged ? <span><Link to="/forecast">Forecast</Link></span> : null }
            </nav>
        );
    }

    return (
        <Router>
            <header className="w-full top-0 fixed flex justify-center md:justify-between items-center h-24 bg-yellow-500 text-black">
                <span className="logo md:ml-8">Landing</span>
                { renderLinks() }
            </header>

            <div className="mt-24">
                <Switch>
                    <Route exact path="/">
                        <Home />
                    </Route>
                    <Route path="/login">
                        { logged ? <Redirect to="/" /> : <Login loginHandler={loginHandler} /> }
                    </Route>
                    <Route path="/forecast">
                        { logged ? <Forecast /> : <Redirect to="/login" /> }
                    </Route>
                    <Route path="*">
                        <Redirect to="/" />
                    </Route>
                </Switch>
            </div>

        </Router>
    );
};

