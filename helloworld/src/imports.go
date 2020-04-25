package main

import (
	_ "github.com/nareshkumarthota/flogocomponents/activity/methodinvoker"
	_ "github.com/project-flogo/contrib/trigger/rest"
	_ "github.com/project-flogo/core/app/propertyresolver"
	_ "github.com/project-flogo/messaging-contrib/kafka/trigger"
	_ "github.com/project-flogo/messaging-contrib/kafka/activity"
	_ "github.com/project-flogo/contrib/activity/kafka"
	_ "github.com/project-flogo/contrib/activity/log"
	_ "github.com/project-flogo/contrib/trigger/kafka"
	_ "github.com/project-flogo/microgateway"
)
