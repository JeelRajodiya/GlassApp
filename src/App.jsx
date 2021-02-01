
import React, { Component } from "react";
import { TouchableOpacity,
        View,
        TextInput,
        StyleSheet,
        Text,
        ScrollView,
        SafeAreaView,
        FlatList,
        ActivityIndicator,
        StatusBar,
        Switch
 } from "react-native";
import styles from "./styles.jsx"



function MainWindow(props){
    const [ipt,changeIpt] = React.useState("");
    const [results,changeResults] = React.useState([])
    const [isloading,changeloading] = React.useState(false);
    const [istyping,changeTyping] = React.useState(false)

    return (
        <SafeAreaView style = {styles.MainWindow}>
        <InputView changeipt = {changeIpt}
                     changeresults={changeResults}
                      ipt={ipt}
                      loadingstate = {[isloading,changeloading]}
                      typingState = {[istyping,changeTyping]}
                      />
         <ResultView results = {results}
                    isloading = {isloading}

                    />
        </SafeAreaView>
    )
}

function Input(props){
    return (
        <TextInput value={props.ipt} onChangeText = {(text)=>{
            let value = text;
            props.changeipt(value);
            var [istyping,changeTyping] = props.typingState
            if (istyping) {clearTimeout(istyping)}
            var timer = setTimeout(() =>{fatchResults(value,props.loadingstate,props.changeresults)},800)
            changeTyping(timer)


        }}
            autoCapitalize="none"
            autoCompleteType="off"
            autoCorrect={false}
            autoFocus={true}
            selectionColor = "#4287f5"
            style = {styles.Input}/>

    )
}

function InputView(props){
    return (
        <View style={styles.InputView}>
            <Input changeipt={props.changeipt} 
                changeresults = {props.changeresults}
                loadingstate={props.loadingstate}  
                typingState = {props.typingState}

                ipt = {props.ipt}/>
            <OkBtn changeipt = {props.changeipt} changeresults = {props.changeresults}/>
        </View>
    )
}

function OkBtn(props){
    const TextStyle = {
        color:"red",
        alignSelf:"center",
        fontWeight:"bold"
    }
    return (
        <TouchableOpacity style={styles.OkBtn}  
            onPress={() => {
                props.changeipt("")
                props.changeresults([])
                }

            } >
        <Text style={TextStyle}>CLEAR</Text>
        </TouchableOpacity>
    )
}

function ResultCard(props){
    return (
        
        <View style = {styles.ResultCard} key = {props.text}>
            <Text style={{color:"#4287ff",fontSize:15}}> {props.text} </Text>
        </View>
    )
}

function ResultView(props){
  const ScrollViewStyle = {
    margin:0,
    padding:0
  }
    return (
        <View style = {styles.ResultView}>
        <StatusBar hidden={false} translucent={true}/>
        <ScrollView style = {ScrollViewStyle}>
            <ActivityIndicator color="#4287f5" size="large" animating={props.isloading} style = {{display: props.isloading ? "flex" : "none",padding:20}}/>    
            <FlatList data = {props.results} renderItem = { ({item}) => <ResultCard text = {item}>  </ResultCard>}/>
        </ScrollView>
        </View>

    )
}


async function fatchResults(input,loadingstate,callback){
    const [isloading,changeloading] = loadingstate;
    
    if (input){

        changeloading(true)
        try{
            var response = await fetch('https://goserve121.herokuapp.com/search = <'+ input+'>' );
            response = await response.json();
            changeloading(false)
            callback(response)
            
        } catch (e) {
            changeloading(false)
            }


    } else {
        callback([])
    }

}

export default function App(){
    return (
        <MainWindow/>
                
    )
}
