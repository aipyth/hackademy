import React, { useState } from 'react';
import Storage from '../Storage';

interface LoginData {
    email: string,
    password: string,
};

export interface LoginProps {
    loginHandler: (logged: boolean) => void,
};
export function Login(props: LoginProps) {
    const submitLoginHandler = (data: LoginData): void => {
        Storage.saveUser(data);
        props.loginHandler(true);
    }

    return (
        <div className="login">
            <h1>Login</h1>
            <LoginForm callback={submitLoginHandler} />
        </div>
    );
};

interface LoginFormProps {
    callback: (data: LoginData) => void,
};
function LoginForm(props: LoginFormProps) {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
        setEmail(event.target.value);
    }

    const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
        setPassword(event.target.value);
    }

    const handleSubmit = (event: React.SyntheticEvent): void => {
        event.preventDefault();
        props.callback({ email, password });
    }

    return (
        <form onSubmit={handleSubmit}>
            <LoginEmailInput
                email={email}
                callback={handleEmailChange}
            />
            <LoginPasswordInput
                password={password}
                callback={handlePasswordChange}
            />
            <button type="submit">Login</button>
        </form>
    ); 

}

interface LoginEmailInputProps {
    email: string,
    callback: React.ChangeEventHandler<HTMLInputElement>,
};
function LoginEmailInput(props: LoginEmailInputProps) {
    return (
        <input type="email" value={props.email} onChange={props.callback} placeholder="Email" />
    );
}

interface LoginPasswordInputProps {
    password: string,
    callback: React.ChangeEventHandler<HTMLInputElement>,
};
function LoginPasswordInput(props: LoginPasswordInputProps) {
    return (
        <input type="password" value={props.password} onChange={props.callback} placeholder="Password" />
    );
}
