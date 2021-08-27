import React from "react";

export interface ForecastProps {
    
};
interface ForecastState {
    search: string,
    data: any,
    timer: number | null,
};

const API_KEY = process.env.REACT_APP_WEATHER_API_KEY;

export class Forecast extends React.Component<ForecastProps, ForecastState> {
    constructor(props: ForecastProps) {
        super(props);
        this.state = {
            search: "",
            data: {},
            timer: null,
        };
        this.handleSearchInputUpdate = this.handleSearchInputUpdate.bind(this);
    }

    handleSearchInputUpdate(e: React.FormEvent<HTMLInputElement>): void {
        if (this.state.timer) { clearTimeout(this.state.timer); }
        const gf = this.getCityWeather.bind(this)
        const t = setTimeout(gf, 1000, e.currentTarget.value)
        this.setState({
            search: e.currentTarget.value,
            timer: t,
        })
    }

    async getCityWeather(city: string) {
        const url = `https://api.openweathermap.org/data/2.5/weather?q=${city}&appid=${API_KEY}`
        console.log("Request url", url);
        const resp = await fetch(url);
        if (resp.ok) {
            const body = await resp.json()
            this.setState({ data: body });
        } else {
            console.log(`Response failed with status code`, resp.status);
        }
    }

    render() {
        return (
            <div className="forecast">
                <CityInput
                    text={this.state.search}
                    callback={this.handleSearchInputUpdate}
                />
                <ForecastTable
                    data={this.state.data}
                />
                
            </div>
        );
    }
};

interface ForecastTableProps {
    data: any,
};
class ForecastTable extends React.Component<ForecastTableProps> {
    renderWeather() {
        if (this.props.data.weather) {
            const items = this.props.data.weather.map((value: any) => {
                return (
                    <span className="weather-item" key={value.id}>{value.main}: {value.description}</span>
                );
            });
            return (
                <td>
                    {items}
                </td>
            );
        }
        return null;
    }

    render() {
        return (
            <table>
                <caption>Forecast table</caption>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Country</th>
                        <th>Weather</th>
                    </tr>

                </thead>
                <tbody>
                    <tr>
                        <td>{this.props.data.name}</td>
                        <td>{this.props.data.sys?.country}</td>
                        {this.renderWeather()}
                    </tr>

                </tbody>
            </table>
        );
    }
};

type CallbackFunction = (e: React.FormEvent<HTMLInputElement>) => void;
interface CityInputProps {
    text: string,
    callback: CallbackFunction,
};
interface CityInputState {

};

class CityInput extends React.Component<CityInputProps, CityInputState> {
    render() {
        return (
            <form>
                <label htmlFor="city-input">City</label>
                <input id="city-input" type="text" value={this.props.text} onChange={this.props.callback} />
            </form>
        );
    }
};
