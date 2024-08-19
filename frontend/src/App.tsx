import {useEffect, useState} from 'react';
import './App.css';
import {RegularLayout} from "./layout/RegularLayout";
import {FundraisingTable} from "./components/FundraisingTable/FundraisingTable";
import {CreateFundraising, DeleteFundraising, GetFundraisingList, SyncFundraising} from "../wailsjs/go/app/App";
import {fundraising} from "../wailsjs/go/models";
import {Button, Input, Space} from "antd";
import FetchListDTO = fundraising.FetchListDTO;
import FundraisingWithHistory = fundraising.FundraisingWithHistory;
import FetchListResponse = fundraising.FetchListResponse;

function App() {
    const [fundraisingUrl, setFundraisingUrl] = useState('');
    const updateURL = (e: any) => setFundraisingUrl(e.target.value);
    const [fundraisings, setFundraisings] = useState<fundraising.FetchListResponse>(new FetchListResponse({data: [], total_pages: 1}));
    const [isLoading, setIsLoading] = useState(false);

    const [sortOrder, setSortOrder] = useState<'ASC' | 'DESC'>('DESC');
    const [sortField, setSortField] = useState<keyof FundraisingWithHistory>('id');
    const [page, setPage] = useState(1);
    const [limit, setLimit] = useState(10);


    const handleGetFundraisingList = () => {
        setIsLoading(true)
        const dto = new FetchListDTO({
            page: page,
            limit: limit,
            sort: {
                column: sortField,
                order: sortOrder
            }
        })
        console.log(dto)
        GetFundraisingList(dto).then((response) => {
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
    }, [page, limit, sortField, sortOrder]);


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

    const handleChangePage = (page: number, size: number) => {
        setPage(page)
        setLimit(size)
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
                <FundraisingTable total={fundraisings.total} handleChangePage={handleChangePage} isLoading={isLoading} handleDeleteFundraising={handleDeleteFundraising} handleSyncFundraising={handleSyncFundraising} items={fundraisings.data}/>
            </RegularLayout>
        </div>
    )
}

export default App
