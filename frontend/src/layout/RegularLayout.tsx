import React from "react";
import {Header} from "../components/Header/Header";

export const RegularLayout = ({children}: { children: React.ReactNode }) => {
    return (
        <div className="layout">
            <Header/>
            <div className="content">
                {children}
            </div>
        </div>
    )
}