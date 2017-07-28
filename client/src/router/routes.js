import Index from '@/components/Index'
import SignUp from "@/components/SignUp"
import Profile from "@/components/Profile"
import NotFound from "@/components/NotFound"
import Logout from "@/components/Logout"

export default [

  {
    path: "/",
    name: "Index",
    component: Index
  },
  {
    path: "/signup",
    name: "SignUp",
    component: SignUp
  },
  {
    path: "/profile",
    name: "Profile",
    component: Profile
  },
  {
    path: "/logout",
    name: "Logout",
    component: Logout
  },
  {
    path: "*",
    name: "NotFound",
    component: NotFound
  }

]
