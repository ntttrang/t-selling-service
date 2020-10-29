# Creating a RESTful API with Golang

## About <a name = "about"></a>

This project practices to build a web application by Golang.

## Prerequisites

- Visual code studio or other IDE
- PlantUML

## Usage <a name = "usage"></a>

<ul>
    <li>
    After download source code, you can run these commands: <br/>
        go mod tidy<br/>
        go mod vendor<br/>
        make execute<br/>
    </li>
    <li>
    Try to call API: <br/>
    [GET] http://localhost:8089/api/v1/admin/users <br/>
    [POST] http://localhost:8089/api/v1/admin/users <br/>
    [PUT] http://localhost:8089/api/v1/admin/users <br/>
    [DELETE] http://localhost:8089/api/v1/admin/users <br/>
    </li>
</ul>


## How to see the system flow
Right-click at `system_flow.puml` and select `Export Workspace Diagrams` <br/>
See the result at `/out/system_flow`
