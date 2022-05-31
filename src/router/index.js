// createWebHistory allows you to use the back arrow in a webpage to return to the previous page
import { createRouter, createWebHistory } from "vue-router";
import MyBody from './../components/MyBody.vue' //importing Mybody component to allow a better way of routing
import MyLogin from '../components/MyLogin.vue' // importing mylogin component to allow a better way of routing
import MyBooks from '../components/MyBooks.vue'
import MyBook from '../components/MyBook.vue'
import BooksAdmin from './../components/BooksAdmin.vue'
import BookEdit from './../components/BookEdit.vue'
import MyUsers from'./../components/MyUsers.vue'
import User from'./../components/UserEdit.vue'
import Security from "@/components/security";

// declaring your routes  path, name, and component it links too.
const routes = [
  {
    path: '/',
    name: 'Home',
    component: MyBody,
  },
  { // go to myheader and see how you linked these onto the navbar links
    path: '/login', //this is allowing you to make a new path to a mylogin page this way you can route new pages made to this one
    name: 'Login',
    component: MyLogin,
  },
  {
    path: "/books",
    name: 'MyBooks',
    component: MyBooks,
  },
  {
    path: "/books/:bookName", // :bookName  purpose of colon is to change depending on book we are looking at
    name: 'MyBook',
    component: MyBook,
  },
  {
    path: "/admin/books",
    name: "BooksAdmin",
    component: BooksAdmin,
  },
  {
    path: "/admin/books/:bookId", // this will change depending on the book
    name: 'BookEdit',
    component: BookEdit,
  },
  {
    path: "/admin/users",
    name: 'MyUsers',
    component: MyUsers,
  },
  {
    path: "/admin/users/:userId",
    name: "MyUser",
    component: User,
  },
]
//  this is where you create the router to be able to call the routes you declared by then exporting it to be used in App.vue 
const router = createRouter({ history: createWebHistory(), routes })
router.beforeEach(() => { // do something before every route is beforeEach
  Security.checkToken(); 
})
export default router 