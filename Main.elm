port module Main exposing (main)

import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Example exposing (..)


port search : SearchRequest -> Cmd msg


port respond : (SearchResponse -> msg) -> Sub msg


type alias Model =
    { names : List String
    }


init : ( Model, Cmd Msg )
init =
    ( { names = []
      }
    , Cmd.none
    )


type Msg
    = Search String
    | Response SearchResponse


view : Model -> Html Msg
view model =
    div []
        [ input [ placeholder "search names...", onInput Search ] []
        , case model.names of
            [] ->
                p [] [ text "no hits" ]

            lst ->
                lst
                    |> List.map (\name -> li [] [ text name ])
                    |> ul []
        ]


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        Search query ->
            ( model, search <| SearchRequest query )

        Response response ->
            ( { model | names = response.items }, Cmd.none )


subscriptions : Model -> Sub Msg
subscriptions model =
    respond Response


main : Program Never Model Msg
main =
    Html.program
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }
