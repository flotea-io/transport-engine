/*
* Project: FI-WARE
* Copyright (c) 2014 Center for Internet Excellence, University of Oulu, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENCE
*/

package utils

import (
  //"fmt"
  //"strconv"
  //"time"
  //"encoding/json"
)

type DataManager struct {
}

func (d *DataManager) Name() string {
    return "TimeSpan"
}

func (d *DataManager) parseShedule(schedule []int) bool{

  //https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/
  /*
  myJsonString := `{"some":"json"}`

  // `&myStoredVariable` is the address of the variable we want to store our
  // parsed data in
  json.Unmarshall([]byte(myJsonString), &myStoredVariable)
  */


  //&$common_params, &$json_struct
  //if (schedule["begin_time"] && schedule["end_time"] && schedule["min_minutes"]){
  //  begin_time = schedule["begin_time"]
  //  end_time = schedule["end_time"]




    /* If schedule given as a search parameter, combine it with POIs schedule
    //using 'and' operator
    if (isset($common_params['schedule']))
    {
        $schedule = array("and" => array($schedule, $common_params['schedule']));
    }

    $res_begintime = array();
    $res_endtime = array();
    $start_event = array($begin_time['year'], $begin_time['month'], $begin_time['day'], $begin_time['hour'], $begin_time['minute'], $begin_time['second']);
    $end_limit = array($end_time['year'], $end_time['month'], $end_time['day'], $end_time['hour'], $end_time['minute'], $end_time['second']);
    $result = $this->timespan->find_open_time($schedule, $common_params['min_minutes']*60, $start_event, $end_limit, $res_begintime, $res_endtime);

    //Filter POIs from $json_struct that do not fulfill the time constraints...
    if ($result == False)
    {
        unset($json_struct["pois"][$uuid]);
    }*/

    //  }
    
    return true
}


/*
public function parseShedule(&$common_params, &$json_struct)
    {
        if (isset($common_params['begin_time']) and isset($common_params['end_time']) and isset($common_params['min_minutes']))
        {

            $begin_time = $common_params['begin_time'];
            $end_time = $common_params['end_time'];

            foreach(array_keys($json_struct["pois"]) as $uuid)
            {

                $fw_time = $this->dbutils->getComponentMongoDB("fw_time", $uuid, false);

                //Remove POI from $json_struct as it does not contain fw_time...
                if ($fw_time == NULL)
                {
                    unset($json_struct["pois"][$uuid]);
                    continue;
                }

                $schedule = $fw_time['schedule'];

                //If schedule given as a search parameter, combine it with POIs schedule
                //using 'and' operator
                if (isset($common_params['schedule']))
                {
                    $schedule = array("and" => array($schedule, $common_params['schedule']));
                }

                $res_begintime = array();
                $res_endtime = array();
                $start_event = array($begin_time['year'], $begin_time['month'], $begin_time['day'], $begin_time['hour'], $begin_time['minute'], $begin_time['second']);
                $end_limit = array($end_time['year'], $end_time['month'], $end_time['day'], $end_time['hour'], $end_time['minute'], $end_time['second']);
                $result = $this->timespan->find_open_time($schedule, $common_params['min_minutes']*60, $start_event, $end_limit, $res_begintime, $res_endtime);

                //Filter POIs from $json_struct that do not fulfill the time constraints...
                if ($result == False)
                {
                    unset($json_struct["pois"][$uuid]);
                }
            }
        }
    }
*/
