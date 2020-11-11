

#import rounds.csv into data

agg = aggregate(list("Time" = data$TIME), list("Model" = data$MODEL, "Clients" = data$CLIENTS, "Node ID" = data$NODE_ID, "Round" = data$ROUND), FUN=sum)

boxplot(Time ~ Clients, data = agg)
