//import axios from "axios"
//import {Card,Header,Form,Input,Icon} from "semantic-ui-react"
//import { useState,useEffect } from "react";
import { Header,Form, Input } from "semantic-ui-react"
import { useState } from "react"

//let endpoint = "http://localhost:9000";

function ToDoList()
{
   const [task,setTask] = useState("")
  //  const[items,setItems] = useState([])

    //useEffect(()=>{


    //},[])

    return (
        <div>
            <div className="row">
                <Header className="header" as="h2" color="yellow">
                    TO DO LIST
                </Header>
            </div>
            <div className="row">
                <Form onSubmit={onSubmit}>
                    <Input
                        type="text"
                        name="task"
                        onChange={setTask}
                        value={task}
                        placeholder="create task"
                        />

                </Form>
            </div>
        </div>

    )

}


export default ToDoList