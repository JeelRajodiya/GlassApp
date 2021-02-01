import {StyleSheet} from "react-native";


var styles = StyleSheet.create({

    MainWindow:{
        paddingTop:35,
        padding:10,
        flex:1,
        backgroundColor:"black"
    },

    Input:{
        flex:1,
        backgroundColor: '#191919',
        color:"#2bd1ff",
        borderRadius:100,
        paddingHorizontal:20,
        borderWidth:1,
        borderColor:"white"
    },
    OkBtn:{
      borderWidth:1,
        borderColor:"#940d03",  
        backgroundColor:"#360404",
        width:60,
        alignContent:"center",
        justifyContent:"center",
        borderRadius:50,
        marginHorizontal:5, 

        
    },
    InputView:{
        flexDirection:"row",

        height:50
    },
    ResultView:{
        flex:20,
        backgroundColor:"#000",
        marginVertical:10,
        borderWidth:1,
        borderColor:"#777",
        borderRadius:8,
        paddingVertical:5
        
        
    },
    ResultCard:{
      backgroundColor:"#1f1f1f",
        shadowColor: 'white',
        borderColor:"#4287f5",
        borderWidth:1,
        borderRadius:15,
        paddingVertical:13,
        padding:17,
        marginVertical:4,
        margin:5
        
    }
})
export default styles