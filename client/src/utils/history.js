import {createHistory} from "history"
import createHashHistory from 'history/lib/createHashHistory'

export default createHashHistory({
    queryKey: false
})
