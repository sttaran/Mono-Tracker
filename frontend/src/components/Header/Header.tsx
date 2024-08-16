import React from "react";
import {Divider, Image, Space} from "antd";
import logo from '../../assets/images/logo-universal.png';

export const Header = () => {
    return (
        <header>
            <Space className="header" align="start" direction="horizontal" wrap>
                <img width={100} src={logo} />
                <h1>Mono Tracker</h1>
            </Space>
            <Divider/>
        </header>


    )
}