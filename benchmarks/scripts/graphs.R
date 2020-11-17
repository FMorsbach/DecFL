

#import rounds.csv into data
agg = aggregate(list("Time" = data$TIME), list("Model" = data$MODEL, "Clients" = data$CLIENTS, "Node ID" = data$NODE_ID, "Round" = data$ROUND), FUN=sum)
boxplot(Time ~ Clients, data = agg)



# import phases.csv into data
library(ggplot2)
data$CLIENTS <- as.factor(data$CLIENTS)
data$STEP <- as.factor(data$STEP)
data$MODEL <- as.factor(data$MODEL)
data$NODE_ID <- as.factor(data$NODE_ID)
ggplot(data, aes(fill=data$STEP, y=data$TIME, x=data$CLIENTS)) + 
  geom_bar(position="fill", stat="identity") +
  xlab("Clients") +
  ylab("")
