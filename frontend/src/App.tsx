import {useEffect, useState} from 'react';
import './App.css';
import {RegularLayout} from "./layout/RegularLayout";
import {FundraisingTable} from "./components/FundraisingTable/FundraisingTable";
import {CreateFundraising, DeleteFundraising, GetFundraisingList, SyncFundraising} from "../wailsjs/go/app/App";
import {fundraising} from "../wailsjs/go/models";
import {Button, Input, Space} from "antd";

function App() {
    const [fundraisingUrl, setFundraisingUrl] = useState('');
    const updateURL = (e: any) => setFundraisingUrl(e.target.value);
    const [fundraisings, setFundraisings] = useState<fundraising.FundraisingWithHistory[]>([]);
    const [isLoading, setIsLoading] = useState(false);

    const handleGetFundraisingList = () => {
        setIsLoading(true)
        GetFundraisingList().then((response) => {
            console.log(response)
            setFundraisings(response)
        }).catch((error) => {
            console.error(error)
        }).finally(() => {
            setIsLoading(false)
        })
    }

    useEffect(() => {
        handleGetFundraisingList()
    }, []);


    function createFundraising() {
        setIsLoading(true)
        CreateFundraising(fundraisingUrl).then((response) => {
            console.log(response)
            alert('Fundraising created. Wait few seconds to sync')
            return response
        }).then((response) => handleSyncFundraising(response, true))
            .catch((error) => {
            alert(`Error creating: ${error}` )
        }).finally(()=> {
            setFundraisingUrl('')
            setIsLoading(false)
        })
    }

    const handleSyncFundraising = async (id: number, initial = false) => {
        setIsLoading(true)
        try {
            try {
                 await SyncFundraising(id, initial);
                handleGetFundraisingList();
                alert('Fundraising synced');
            } catch (error) {
                alert(`Error syncing: ${error}`);
            }
        } finally {
            setIsLoading(false);
        }

    }

    const handleDeleteFundraising = async (id: number) => {
        setIsLoading(true)
        try {
            try {
                await DeleteFundraising(id);
                handleGetFundraisingList();
                alert('Fundraising deleted');
            } catch (error) {
                alert(`Error deleting: ${error}`);
            }
        } finally {
            setIsLoading(false);
        }
    }

    return (
        <div id="App">
            <RegularLayout>
                <p id="title">Please enter fundraising URL</p>
                <div id="input" className="input-box">
                    <Space>
                        <Input disabled={isLoading} id="name" value={fundraisingUrl} onChange={updateURL} autoComplete="off" name="input"
                               type="text"/>
                        <Button disabled={isLoading} type="primary" onClick={createFundraising}>Add Fundraising</Button>
                    </Space>
                    <br/>
                    <br/>
                    </div>
                <FundraisingTable isLoading={isLoading} handleDeleteFundraising={handleDeleteFundraising} handleSyncFundraising={handleSyncFundraising} items={fundraisings}/>
            </RegularLayout>
        </div>
    )
}

export default App
