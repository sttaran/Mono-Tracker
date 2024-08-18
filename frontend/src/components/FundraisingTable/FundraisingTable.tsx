import React from "react";
import {fundraising} from "../../../wailsjs/go/models";
import {Button, Popover, Space, Table, Typography} from "antd";
import relativeTime from 'dayjs/plugin/relativeTime'
import dayjs from "dayjs";

dayjs.extend(relativeTime);

interface FundraisingTableProps {
    items: fundraising.FundraisingWithHistory[]
    handleSyncFundraising: (id: number) => void
    handleDeleteFundraising: (id: number) => void
    isLoading: boolean
}

function getColorForPercent(percent: number) {
    if (percent < 25) {
        return 'red'
    } else if (percent < 50) {
        return 'orange'
    } else if (percent < 75) {
        return 'yellow'
    } else {
        return 'green'
    }
}

function getColorForRaised(raised: number, previousRaised: number) {
    switch (true) {
        case raised > previousRaised:
            return 'green'
        case raised < previousRaised:
            return 'red'
        default:
            return 'gray'
    }
}

export const FundraisingTable: React.FC<FundraisingTableProps> = ({items, isLoading, handleSyncFundraising, handleDeleteFundraising}) => {
    return (
        <Table bordered loading={isLoading} rowKey="id" dataSource={items} columns={[
            {
                title: 'Name',
                dataIndex: 'name',
                key: 'name',
            },
            {
                title: 'Description',
                dataIndex: 'description',
                key: 'description',
                width: 400,
            },

            {
                title: 'Goal',
                dataIndex: 'goal',
                key: 'goal',
                render: (text: number) => <span>{new Intl.NumberFormat("ua-UA").format(text)}</span>
            },
            {
                title: 'Raised %',
                dataIndex: 'raised',
                key: 'raised_percent',
                render: (_: number, record) => {
                    const value = record.goal && record.history[0]?.raised ?  Math.round((record.history[0]?.raised/ record.goal) * 100) : 0
                    return <span style={{
                        color: getColorForPercent(value)
                    }}>{value}%</span>
                }
            },
            {
                title: 'Raised',
                dataIndex: 'raised',
                key: 'raised',
                render: (text: number, record) => {
                     const raised = record.history[0]?.raised
                    const previousRaised = record.history[1]?.raised
                    console.log(raised, previousRaised)
                    return <div>
                        <span>{raised ? new Intl.NumberFormat("ua-UA").format(raised) : "unknown"}</span>
                        <br/>
                        <span style={{color: getColorForRaised(raised, previousRaised)}}>
                            {raised && previousRaised && raised > previousRaised && '+' }
                            {raised && previousRaised && raised < previousRaised && '-' }
                            {raised && previousRaised ? `${new Intl.NumberFormat('ua-UA').format(Math.abs(raised - previousRaised))}` : ''}
                        </span>
                    </div>
                }
            },
            {
                title: 'Last synced',
                dataIndex: 'history',
                key: 'last_synced',
                render: (data: fundraising.FundraisingWithHistory["history"]) => {
                    return <span>{data[0]?.sync_time ? dayjs().to(dayjs(data[0].sync_time)) : 'never'}</span>
                }
            },
            {
                title: 'History',
                dataIndex: 'history',
                key: 'history',
                render: (item: fundraising.FundraisingWithHistory["history"]) => {
                    const options = {
                        year: "numeric",
                        month: "numeric",
                        day: "numeric",
                        hour: "numeric",
                        minute: "numeric",
                        second: "numeric",
                        hour12: false,
                    } as const

                    return <div>
                        {item.map((history) => <div key={history.id}>
                        <span style={{
                            color: 'gray'
                        }}>{new Intl.DateTimeFormat("eu-DE",  options).format(new Date(history.sync_time))}{" "}</span>
                            <span style={{
                                color: 'green'
                            }}>{new Intl.NumberFormat("ua-UA").format(history.raised)}</span>
                        </div>)}
                    </div>
                }
            },
            {
                title: 'URL',
                dataIndex: 'url',
                key: 'url',
                render: (text: string) => <a href={text}>{text}</a>
            },
            {
                title: 'Action',
                key: 'action',
                render: (text: string, record) => (
                    <Space>
                        <Button onClick={()=>handleSyncFundraising(record.id)}>
                            Sync
                        </Button>
                        <Popover trigger="click" content={
                            <Space direction="vertical" align="center">
                                <Typography.Paragraph>
                                    Fundraising will be deleted permanently.
                                </Typography.Paragraph>

                                <Button onClick={()=>handleDeleteFundraising(record.id)} danger type="primary">
                                    DELETE
                                </Button>
                            </Space>
                        }>
                            <Button  type="dashed" danger>
                                Delete
                            </Button>
                        </Popover>
                    </Space>
                ),
            }
        ]}/>
    )
}