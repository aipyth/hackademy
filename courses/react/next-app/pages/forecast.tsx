import Layout from '../components/Layout';
import React, { useState } from "react";

const API_KEY = process.env.REACT_APP_WEATHER_API_KEY;

export default function Forecast() {
    const [search, setSearch] = useState("");
    const [data, setData] = useState({});
    const [timer, setTimer] = useState(null);


    const getCityWeather = async (city: string) => {
        const url = `https://api.openweathermap.org/data/2.5/weather?q=${city}&appid=${API_KEY}`
        const resp = await fetch(url);
        if (resp.ok) {
            const body = await resp.json()
            setData(body);
        } else {
            console.log(`Response failed with status code`, resp.status);
        }
    }

    const handleSearchInputUpdate = (e: React.FormEvent<HTMLInputElement>): void => {
        if (timer) { clearTimeout(timer); }
        const t = setTimeout(getCityWeather, 1000, e.currentTarget.value)
        setSearch(e.currentTarget.value);
        setTimer(t);
    }

    return (
        <Layout title="Forecast Next-App">
        <div className="forecast">
            <CityInput
                text={search}
                callback={handleSearchInputUpdate}
            />
            <ForecastTable
                data={data}
            />
            
        </div>
        </Layout>
    );
};

interface ForecastTableProps {
    data: any,
};
function ForecastTable(props: ForecastTableProps) {
    const renderWeather = () => {
        if (props.data.weather) {
            const items = props.data.weather.map((value: any) => {
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
    };

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
                    <td>{props.data.name}</td>
                    <td>{props.data.sys?.country}</td>
                    {renderWeather()}
                </tr>

            </tbody>
        </table>
    );
}

interface CityInputProps {
    text: string,
    callback: (e: React.FormEvent<HTMLInputElement>) => void,
};
function CityInput(props: CityInputProps) {
    return (
        <form>
            <label htmlFor="city-input">City</label>
            <input id="city-input" type="text" value={props.text} onChange={props.callback} />
        </form>
    );
}

