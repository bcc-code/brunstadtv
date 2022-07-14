import {RouteRecordRaw} from "vue-router"

export default [
    {
        path: "/",
        name: "home",
        component: () => import("@/layout/Home.vue"),
        children: [
            {
                path: "settings",
                name: "settings",
                component: () => import("@/views/Settings.vue")
            }
        ]
    }
] as RouteRecordRaw[]