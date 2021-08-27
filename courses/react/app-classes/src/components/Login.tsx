import React from 'react';
import Storage from '../Storage';

interface LoginData {
    email: string,
    password: string,
};

export interface LoginProps {
    loginHandler: (logged: boolean) => void,
};
interface LoginState {
};
export class Login extends React.Component<LoginProps, LoginState> {
    constructor(props: LoginProps) {
        super(props);
        this.submitLoginHandler = this.submitLoginHandler.bind(this);
    }

    submitLoginHandler(data: LoginData): void {
        Storage.saveUser(data);
        this.props.loginHandler(true);
    }

    render() {
        return (
            <div className="login">
                <h1>Login</h1>
                <LoginForm callback={this.submitLoginHandler} />
            </div>
        );
    }
};

interface LoginFormProps {
    callback: (data: LoginData) => void,
};
interface LoginFormState {
    email: string,
    password: string,
};
class LoginForm extends React.Component<LoginFormProps, LoginFormState> {
    constructor(props: LoginFormProps) {
        super(props);
        this.state = {
            email: "",
            password: "",
        };
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleEmailChange = this.handleEmailChange.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
    }

    handleSubmit(event: React.SyntheticEvent): void {
        event.preventDefault();
        this.props.callback({
            email: this.state.email,
            password: this.state.password,
        });
    }

    handleEmailChange(event: React.ChangeEvent<HTMLInputElement>): void {
        this.setState({ email: event.target.value });
    }

    handlePasswordChange(event: React.ChangeEvent<HTMLInputElement>): void {
        this.setState({ password: event.target.value });
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <LoginEmailInput
                    email={this.state.email}
                    callback={this.handleEmailChange}
                />
                <LoginPasswordInput
                    password={this.state.password}
                    callback={this.handlePasswordChange}
                />
                <button type="submit">Login</button>
            </form>
        );
    }
};

interface LoginEmailInputProps {
    email: string,
    callback: React.ChangeEventHandler<HTMLInputElement>,
};
class LoginEmailInput extends React.Component<LoginEmailInputProps> {
    render() {
        return (
            <input type="email" value={this.props.email} onChange={this.props.callback} placeholder="Email" />
        );
    }
};

interface LoginPasswordInputProps {
    password: string,
    callback: React.ChangeEventHandler<HTMLInputElement>,
};
class LoginPasswordInput extends React.Component<LoginPasswordInputProps> {
    render() {
        return (
            <input type="password" value={this.props.password} onChange={this.props.callback} placeholder="Password" />
        );
    }
}
