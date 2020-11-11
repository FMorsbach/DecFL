

#import rounds.csv into data
agg = aggregate(list("Time" = data$TIME), list("Model" = data$MODEL, "Clients" = data$CLIENTS, "Node ID" = data$NODE_ID, "Round" = data$ROUND), FUN=sum)
boxplot(Time ~ Clients, data = agg)



# import phases.csv into data
data$CLIENTS <- as.factor(b$CLIENTS)
data$STEP <- as.factor(b$STEP)
data$MODEL <- as.factor(b$MODEL)
data$NODE_ID <- as.factor(b$NODE_ID)
ggplot(data, aes(fill=data$STEP, y=data$TIME, x=data$CLIENTS)) + 
  geom_bar(position="fill", stat="identity") +
  xlab("Clients") +
  ylab("")
